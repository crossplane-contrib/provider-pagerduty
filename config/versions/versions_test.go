package versions_test

import (
	"testing"

	"github.com/crossplane/upjet/v2/pkg/apitesting/roundtrip"
	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/utils/ptr"

	apisCluster "github.com/crossplane-contrib/provider-pagerduty/apis/cluster"
	alertv1alpha1 "github.com/crossplane-contrib/provider-pagerduty/apis/cluster/alert/v1alpha1"
	alertv1alpha2 "github.com/crossplane-contrib/provider-pagerduty/apis/cluster/alert/v1alpha2"
	apisNamespaced "github.com/crossplane-contrib/provider-pagerduty/apis/namespaced"
	"github.com/crossplane-contrib/provider-pagerduty/config"
)

// embeddedKinds are the kinds configured with versions.EmbedSingletonLists.
var embeddedKinds = []schema.GroupKind{
	{Group: "alert.pagerduty.crossplane.io", Kind: "GroupingSetting"},
	{Group: "alert.pagerduty.m.crossplane.io", Kind: "GroupingSetting"},
	{Group: "jira.pagerduty.crossplane.io", Kind: "CloudAccountMappingRule"},
	{Group: "jira.pagerduty.m.crossplane.io", Kind: "CloudAccountMappingRule"},
	{Group: "user.pagerduty.crossplane.io", Kind: "NotificationRule"},
	{Group: "user.pagerduty.m.crossplane.io", Kind: "NotificationRule"},
}

func newScheme(t *testing.T) *runtime.Scheme {
	t.Helper()
	s := runtime.NewScheme()
	if err := apisCluster.AddToScheme(s); err != nil {
		t.Fatal(err)
	}
	if err := apisNamespaced.AddToScheme(s); err != nil {
		t.Fatal(err)
	}
	return s
}

func TestConversionRoundTrip(t *testing.T) {
	rt, err := roundtrip.NewRoundTripTest(config.GetProvider(), config.GetProviderNamespaced(), newScheme(t),
		roundtrip.WithIncludeGroupKinds(embeddedKinds...),
		// The singleton lists hold at most one element.
		roundtrip.WithFuzzerConfig(roundtrip.FuzzerIterations(50), roundtrip.FuzzerNumElements(0, 1)))
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Serialization", rt.TestSerializationRoundtrip)
	t.Run("Conversion", rt.TestConversionRoundtrip)
	// NewRoundTripTest registered the conversions.
	t.Run("GroupingSetting", testGroupingSettingConversion)
}

func testGroupingSettingConversion(t *testing.T) {
	spoke := &alertv1alpha1.GroupingSetting{}
	spoke.SetGroupVersionKind(alertv1alpha1.GroupingSetting_GroupVersionKind)
	spoke.Spec.ForProvider.Config = []alertv1alpha1.ConfigParameters{{
		Aggregate:  ptr.To("all"),
		TimeWindow: ptr.To[float64](300),
	}}
	spoke.Status.AtProvider.Config = []alertv1alpha1.ConfigObservation{{Aggregate: ptr.To("all")}}

	hub := &alertv1alpha2.GroupingSetting{}
	hub.SetGroupVersionKind(alertv1alpha2.GroupingSetting_GroupVersionKind)
	if err := spoke.ConvertTo(hub); err != nil {
		t.Fatal(err)
	}
	wantSpec := &alertv1alpha2.ConfigParameters{Aggregate: ptr.To("all"), TimeWindow: ptr.To[float64](300)}
	if diff := cmp.Diff(wantSpec, hub.Spec.ForProvider.Config); diff != "" {
		t.Errorf("spec.forProvider.config: -want, +got:\n%s", diff)
	}
	if diff := cmp.Diff(&alertv1alpha2.ConfigObservation{Aggregate: ptr.To("all")}, hub.Status.AtProvider.Config); diff != "" {
		t.Errorf("status.atProvider.config: -want, +got:\n%s", diff)
	}

	back := &alertv1alpha1.GroupingSetting{}
	back.SetGroupVersionKind(alertv1alpha1.GroupingSetting_GroupVersionKind)
	if err := back.ConvertFrom(hub); err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(spoke.Spec, back.Spec); diff != "" {
		t.Errorf("v1alpha1 -> v1alpha2 -> v1alpha1 spec: -want, +got:\n%s", diff)
	}
}
