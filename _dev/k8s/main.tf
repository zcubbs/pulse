terraform {
    required_providers {
        kubectl = {
            source  = "gavinbunney/kubectl"
            version = ">= 1.7.0"
        }

        kubernetes-alpha = {
            source  = "hashicorp/kubernetes-alpha"
            version = "0.6.0"
        }
    }
}

provider "kubectl" {
    config_path    = "~/.kube/config"
    config_context = "rancher-desktop"
}

provider "kubernetes" {
    config_path    = "~/.kube/config"
    config_context = "rancher-desktop"
}

provider "kubernetes-alpha" {
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

resource "kubernetes_manifest" "argocd-ingress-route" {
    depends_on = [
        helm_release.argocd
    ]

    manifest = {
        apiVersion = "traefik.containo.us/v1alpha1"
        kind       = "IngressRoute"
        metadata   = {
            name      = "argocd"
            namespace = "argocd"
        }
        spec = {
            entryPoints = [
                "web",
                "websecure"
            ]
            routes = [
                {
                    match    = "Host(`argo.localhost`)"
                    kind     = "Rule"
                    services = [
                        {
                            name = "argocd-server"
                            port = "80"
                        }
                    ]
                }
            ]
        }
    }
}