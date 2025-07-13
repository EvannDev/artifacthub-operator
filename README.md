# 📦 artifacthub-operator

**artifacthub-operator** est un Kubernetes Operator basé sur Helm, développé avec [Operator SDK](https://sdk.operatorframework.io/), permettant de déployer et gérer des instances d’[Artifact Hub](https://artifacthub.io/) de manière déclarative.

> 🚧 **Statut** : Projet en phase initiale / POC — contributions et feedback bienvenus !

---

## 🎯 Objectif

Ce projet vise à fournir un operator natif Kubernetes pour automatiser le déploiement d'Artifact Hub via une ressource personnalisée (CRD), facilitant ainsi son intégration dans des environnements GitOps, CI/CD ou multi-clusters.

---

## ⚙️ Stack technique

- Operator SDK (Helm plugin)
- Helm chart (officiel ou custom pour Artifact Hub)
- Kubernetes ≥ 1.22

---

## 🚀 Installation

```bash
git clone https://github.com/EvannDev/artifacthub-operator.git
cd artifacthub-operator

# Build et déploiement via operator-sdk
make install
make deploy
```

> [!WARNING] 
> ⚠️ Prérequis : kubectl, kustomize, operator-sdk, accès à un cluster avec permissions d’écriture sur les CRDs.

## 📄 Exemple de Custom Resource (CR)

```yaml
apiVersion: artifacthub.evann.dev/v1alpha1
kind: ArtifactHub
metadata:
  name: example
spec:
  # Valeurs spécifiques au chart Helm, par exemple :
  replicaCount: 1
  ingress:
    enabled: true
    hosts:
      - host: artifacthub.local
```

## 🔍 Fonctionnement

Une fois le CR appliqué, l’operator :

* Installe Artifact Hub via Helm dans le namespace cible

* Suit le cycle de vie du CR pour appliquer les mises à jour ou suppressions

* Peut être intégré dans un flux GitOps (Argo CD, Flux, etc.)

## 🧪 Tests

```bash
make test
# Ou via kind ou minikube
```

## 🤝 Contribution

Les PRs sont les bienvenues ! Pour contribuer :

* Fork le repo

* Crée une branche (feat/ma-fonctionnalite)

* Teste localement avec kind ou minikube

* Soumets une PR bien documentée
