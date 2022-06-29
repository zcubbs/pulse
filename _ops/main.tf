terraform {
    required_providers {
        kubectl = {
            source  = "gavinbunney/kubectl"
            version = ">= 1.7.0"
        }
    }
}

provider "kubectl" {
    config_path    = "~/.kube/config"
    config_context = "rancher-desktop"
}

provider "helm" {
    kubernetes {
        config_path    = "~/.kube/config"
        config_context = "rancher-desktop"
    }
}

resource "helm_release" "argocd" {
    name = "argocd"

    repository       = "https://argoproj.github.io/argo-helm"
    chart            = "argo-cd"
    namespace        = "argocd"
    version          = "4.9.3"
    create_namespace = true
    verify           = false

    values = [
        file("argocd/values.yaml")
    ]
}