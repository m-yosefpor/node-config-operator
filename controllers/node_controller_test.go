package controllers

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	configv1alpha1 "github.com/snapp-incubator/node-config-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("nodeMatchNodeConfig", func() {
	It("matches node names with regex", func() {
		node := corev1.Node{ObjectMeta: metav1.ObjectMeta{Name: "worker-1"}}
		match := nodeMatchNodeConfig(node, configv1alpha1.Match{NodeNamePatterns: []string{"worker-.*"}})
		Expect(match).To(BeTrue())
	})

	It("returns false for non matching names", func() {
		node := corev1.Node{ObjectMeta: metav1.ObjectMeta{Name: "master-1"}}
		match := nodeMatchNodeConfig(node, configv1alpha1.Match{NodeNamePatterns: []string{"worker-.*"}})
		Expect(match).To(BeFalse())
	})
})

var _ = Describe("nodeMergeNodeConfig", func() {
	It("updates node labels when they differ", func() {
		node := corev1.Node{ObjectMeta: metav1.ObjectMeta{Name: "node1", Labels: map[string]string{"foo": "bar"}}}
		merge := configv1alpha1.Merge{Labels: map[string]string{"foo": "baz"}}
		updated, match := nodeMergeNodeConfig(node, merge)
		Expect(match).To(BeFalse())
		Expect(updated.Labels["foo"]).To(Equal("baz"))
	})

	It("returns match true when node already has desired labels", func() {
		node := corev1.Node{ObjectMeta: metav1.ObjectMeta{Name: "node1", Labels: map[string]string{"foo": "bar"}}}
		merge := configv1alpha1.Merge{Labels: map[string]string{"foo": "bar"}}
		updated, match := nodeMergeNodeConfig(node, merge)
		Expect(match).To(BeTrue())
		Expect(updated.Labels["foo"]).To(Equal("bar"))
	})
})
