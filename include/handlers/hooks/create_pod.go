package hooks

import corev1 "k8s.io/api/core/v1"

func hookOnPodCreate(pod *corev1.Pod) {
	if pod.Labels == nil {
		pod.Labels = make(map[string]string)
	}

	if pod.Finalizers == nil {
		pod.Finalizers = make([]string, 0)
	}

	pod.Finalizers = append(pod.Finalizers, "flap-operator")
	pod.Labels["hook_on_pod"] = "OK"
}
