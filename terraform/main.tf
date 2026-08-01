provider "kubernetes" {
  config_path = "~/.kube/config"
}

resource "kubernetes_deployment" "go-rest_app" {
  metadata {
    name = "go-rest-app"
    labels = {
      app = "go-rest"
    }
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "go-rest"
      }
    }

    template {
      metadata {
        labels = {
          app = "go-rest"
        }
      }

      spec {
        container {
          name  = "go-rest-container"
          image = "ghcr.io/letsdevapps/go-rest:latest"

          image_pull_policy = "Always"

          port {
            container_port = 8080
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "go-rest_service" {
  metadata {
    name = "go-rest-service"
  }

  spec {
    selector = {
      app = "go-rest"
    }

    port {
      port        = 8080
      target_port = 8080
    }

    type = "NodePort"
  }
}

resource "kubernetes_ingress_v1" "app_ingress" {
  metadata {
    name = "app-ingress"

    annotations = {
      "nginx.ingress.kubernetes.io/rewrite-target" = "/"
    }
  }

  spec {
    ingress_class_name = "nginx"

    rule {
      http {
        path {
          path      = "/go-rest"
          path_type = "Prefix"

          backend {
            service {
              name = kubernetes_service.go-rest_service.metadata[0].name
              port {
                number = 8080
              }
            }
          }
        }
      }
    }
  }
}
