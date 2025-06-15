# ArtifactHub Operator for Kubernetes !

**Repo :** [EvannDev/artifacthub-operator](https://github.com/EvannDev/artifacthub-operator)  
**Description :** Un Kubernetes Operator pour déployer et gérer [Artifact Hub](https://artifacthub.io/) de la CNCF directement depuis votre cluster.  
Basé sur [Operator SDK](https://sdk.operatorframework.io/) et écrit en **Go**.

---

## 🚀 Fonctionnalités

✅ Déploiement automatisé d’une instance Artifact Hub  
✅ Gestion déclarative via CRD `ArtifactHub`  
✅ Synchronisation du `status` pour connaître l’état du déploiement  
✅ Évolutif : support prévu pour Repository, Organization, Roles...

---

## 📂 Structure du projet

```
.
├── api/                # Types API (CRD)
├── config/             # Manifests Kubernetes (RBAC, CRD, Samples)
├── controllers/        # Logique du contrôleur
├── Dockerfile          # Image container
├── Makefile            # Commandes utiles (build, deploy, run)
└── README.md           # Vous êtes ici !
```

---

## ⚙️ Pré-requis

- Un cluster Kubernetes (Minikube, Kind, K3s ou cloud)
- [Operator SDK](https://sdk.operatorframework.io/docs/installation/)
- Docker
- (Optionnel) Un registre container (ex: **GHCR**)

---

## 🔧 Installation

### 1️⃣ Cloner le projet

```bash
git clone https://github.com/EvannDev/artifacthub-operator.git
cd artifacthub-operator
```

### 2️⃣ Builder & pousser l'image

```bash
make docker-build IMG=ghcr.io/<votre-user>/artifacthub-operator:v0.1.0
docker push ghcr.io/<votre-user>/artifacthub-operator:v0.1.0
```

### 3️⃣ Déployer sur le cluster

```bash
make install
make deploy IMG=ghcr.io/<votre-user>/artifacthub-operator:v0.1.0
```

---

## 📜 Créer une ressource ArtifactHub

Un exemple est fourni dans `config/samples/`.

```bash
kubectl apply -f config/samples/core_v1alpha1_artifacthub.yaml
```

---

## ✅ Vérifier

```bash
kubectl get artifacthubs.core.artifacthub.io
kubectl get deployments
kubectl get pods
```

Le champ `status` doit indiquer que l’instance est prête 🎉

---

## 🚧 Roadmap

- [ ] Support pour `Repository`
- [ ] Support pour `Organization` & `Roles`
- [ ] Helm chart pour déploiement simplifié
- [ ] Publication sur OperatorHub

---

## 🤝 Contribuer

PR, Issues et suggestions bienvenues !  
Forkez le repo et proposez vos améliorations 🚀

---

## ⚖️ Licence

[Apache License 2.0](LICENSE)