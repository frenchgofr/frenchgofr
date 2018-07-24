---
title: "Interview de Taichi Nakashima de Mercari"
description: "Interview de Taichi Nakashima de Mercari"
date: 2018-06-20T17:00:00+01:00
draft: false
tags : [
    "interview",
    "deeeet",
]
categories : [
    "interview",
]
---

# Taichi Nakashima de chez Mercari

The [english version is here](https://lemag.sfeir.com/interview-taichi-nakashima-from-mercari/) !

![](/img/post/180620_interview_deeeet/deeeet.jpg)

## Qui es-tu ?

Je m'appelle Taichi Nakashima. Sur twitter je suis [deeeet](https://twitter.com/deeeet) et tu peux m'appeler deeeet !

Je travaille chez [Mercari](http://www.mercari.com), un site de vente entre particuliers; nous laissons les utilisateurs vendre et acheter n'importe quoi, de n'importe où, en quelques secondes depuis leur téléphone. Notre service est disponible au Japon, aux États-Unis et au Royaume-Uni.

Je travaille actuellement dans l'équipe SRE, et je m'occupe principalement du projet de migration vers l'architecture en microservices.

## Que signifie «SRE» ?

Ça correspond à «Site Reliability Engineering». Ça vient de chez Google, le but principal de l'équipe SRE était de permettre aux sites de fonctionner de façon simple et efficace, et de façon plus fiable grâce à leur ingénierie logicielle. Tu peux en découvrir plus avec [leur livre](https://landing.google.com/sre/book.html). 

Notre équipe était anciennement nommé l'équipe d'infrastructure, mais nous l'avons renommée pour suivre les recommandations de Google pour faire de notre produit un service fiable.

## Tu peux décrire ton équipe ?

Actuellemnt Mercari tourne sur une API monolithique et nous essayons de la migrer vers des microservices. Le rôle de [notre équipe](https://open.talentio.com/1/c/mercari/requisitions/detail/7877) est de construire une plateforme qui permette aux développeurs de faire tourner leurs services simplement et de façon fiable.

## Pourquoi cette migration vers les microservices ?

Nous nous développons rapidement et embauchons beaucoup d'ingénieurs. Pour l'instant nous sommes 150 et avons l'ambition d'aller à plus de 1000 ! Notre API monolithique devient trop grosse et complexe. Il est difficile de comprendre et déployer de nouvelles fonctionnalités, et ce rapidement. Nous devons de plus passer beaucoup de temps à communiquer pour trouver notre voie dans un code monolithique.

Si nous pouvons découper les choses en microservices, nous pouvons créer des équipes qui s'occupent d'une partie spécifique du code. En créant de petites équipes qui s'occupent de microservices, nous avons de meilleurs canaux de communications et sommes plus efficaces. Aussi, nous pouvons déployer les services indépendamment et proposer de nouvelles foncitonnalités rapidement.

## Et essayer de ne pas tout casser

Oui ! C'est aussi une des motivations.

## Quels sont les produits que vous proposez ?

Il y a Mercari, notre produit principal. Nous avons une filiale nommée [Souzoh](https://www.souzoh.com/jp/) avec différents produits.

Nous avons aussi démarré une filiale nommée [Merpay](https://www.merpay.com/) qui est un service de paiement.

## Êtes-vous tous sur la même architecture ?

Mmmm non. À Mercari nous avons l'équipe SRE qui s'occupe de l'infrastructure. Chez Souzoh nous n'avons pas de personne en charge de la SRE, uniquement des développeurs d'applications. Ils utilisent [Google App Engine](https://cloud.google.com/appengine/) donc n'ont pas besoin de s'occuper de l'infrastructure.

## Il y a le logo de Google Cloud Platform sur tes slides, vous utilisez quoi ?

Pour la plateforme microservices, nous utilisons GKE, [Google Kubernetes Engine](https://cloud.google.com/kubernetes-engine/), un service managé pour [Kubernetes](https://kubernetes.io/).

## Et où se trouve l'«ancien» Mercari ?

Il est déployé chez un fournisseur d'hébergement japonais.

## Quelle est la roadmap pour la migration ?

Aujourd'hui je travaille sur la création d'une passerelle d'API, je vais te montrer avec [cette slide](https://speakerdeck.com/tcnksm/microservices-on-gke-at-mercari?slide=52):

![](/img/post/180620_interview_deeeet/mercari-microservices.png)


En gros l'architecture actuelle de Mercari est très simple. Nous avons un [nginx](http://nginx.org/en/) tout au dessus, puis l'API applicatif en php et une base de données Mysql. Nous essayons de découper les choses.

Précédemment toutes les requêtes allaient directement sur l'API monolithique. Nous avons créé une passerelle d'API à mettre en amont. Toutes les requêtes arrivent sur celle-ci avant d'être redirigées sur l'API Mercari. Ça ne change rien. Quand nous créons de nouveaux microservices, nous pouvons changer la route au niveau de la passerelle pour aller vers ceux-ci. Pour les services existants, par exemple le service de recherche, si nous souhaitons le migrer, nous allons créer le nouveau service et lorsqu'il sera prêt, changer la route dans la passerelle vers ce nouveau service. C'est notre stratégie de migration.

Cette passerelle et l'architecture qui l'entourent sont très intéressants. Les communications entre l'application mobile et la passerelle sont faits en utilisant les [protocol buffers](https://developers.google.com/protocol-buffers/) via les requêtes HTTP. Entre la passerelle et les microservices nous utilisons [gRPC](https://grpc.io/). C'est très performant.

La passerelle d'API ainsi que les nouveaux microservices sont déployés sur GKE.

Aujourd'hui il est un peu difficile de déployer de nouveaux microservices sur GKE car il fournit uniquement de l'orchestration de conteneurs et les ingénieurs n'ont pas forcément de connaissance de Kubernetes. Nous développons donc des outils pour leur faciliter la vie.

Une autre tache est de migrer tout sur des microservices donc nous avons besoin de réfléchir à une stratégie et à bien planifier car le monolith utilise une seul base de données. Pour les parties en microservices, chaque microservice pourrait être accompagner de sa propre base de données donc nous devons réfléchir à comment découper et répartir nos données. Je pense que ça va prendre beaucoup de temps.

## À propos du livre, tu peux me raconter son histoire ?

[Le livre dont deeet est un des auteurs:](https://www.amazon.co.jp/dp/477418392X)

![](/img/post/180620_interview_deeeet/minna-no-go.jpg)


Nous sommes différents auteurs qui travaillons dans des entreprises différentes. Je pense qu'une personne a reçu une demande d'un éditeur pour écrire un livre sur go, et cette personne a cherché quelques Gophers célèbres au Japon. Nous avons un canal Slack et nous avons discuté pour avoir chacun une partie sur laquelle nous étions experts.

À cette époque je faisais beaucoup d'outils en ligne de commande donc j'ai écrit une partie sur ça. Quelqu'un était très bon sur les tests et a écrit une partie là dessus.

## Ok ! Vas-tu écrire un nouveau livre un jour ?

Pour l'instant je n'ai pas de plan pour ça. La passerelle d'API et les microservices sont écrits en Go... Je ne suis pas certain, mais peut être que dans le futur, si j'en ai la possibilité, j'écrirais un livre sur ce sujet !

## Je reviens sur ton projet, est-ce spécifique à Mercari ou est-ce que ce sera accessible publiquement ?

Pour l'instant c'est spécifique à nos besoins. Je pense que si nous pouvons abstraire certaines choses nous pourrons le rendre open source dans le futur.

## Tu connais bien Go, qu'est-ce que tu préfères dans ce langage ?

Ce sont sa simplicité et sa clarté.

La spécification du langage Go est très simple et le langage nous oblige à écrire beaucoup de choses. Par exemple, en Ruby ou dans d'autres langages, beaucoup de choses sont abstraites et cela rend difficile la lecture et la compréhension de ce qui se passe.

Donc même quand une nouvelle personne rejoint une équipe, il est aisé pour elle de savoir ce qui se passe: il suffit de lire. Go est très bien pour développer en équipe.

## Qu'y-a-t'il de déroutant à propos de Go ?

Je ne sais pas... Je n'ai jamais trouvé quelque chose de déroutant jusqu'à maintenant.

## Quelles sont tes attentes pour Go dans le futur ? Pour Go 2 ?

Aujourd'hui une bonne pratique est d'utiliser le [context](https://golang.org/pkg/context/). Tu devrais l'utiliser d'une certaine manière quand tu crées une goroutine et que tu veux l'annuler. Ça rend le code un peu compliqué et difficile d'accès. J'espère que ce sera bientôt plus simple à utiliser.

Beaucoup de personnes veulent voir les [generics](https://en.wikipedia.org/wiki/Generic_programming) dans Go 2 mais je n'en ai jamais ressenti le besoin. Je n'ai pas été bloqué par l'absence des generics. Je pense que le fait d'introduire les generics rendra le code plus compliqué quand tu dois rentrer dans un code existant.

## Pourquoi le choix de Go à Mercari ?

Car nous avons besoin de performance et de simplicité.

Les gens à Mercari avaient déjà commencé à utiliser Go avant que j'arrive. L'API principale est en php et ce n'est pas forcément assez performant donc nous avons commencé à créer des middlewares en Go pour certaines taches qui nécessitent de meilleurs performances.

## Et pourquoi le choix de Kubernetes ?

Quand nous avons démarré les microservices, nous avons décidé que l'unité à déployer devait être un conteneur. Dans le monde des microservices, il est simple d'être polyglotte et nous souhaitons permettre aux différentes équipes de choisir ce qui est bon pour chaque tache. Si notre unité est le conteneur, il est facile de contrôler et gérer les ressources du point de vue infrastructure. C'est pourquoi les conteneurs sont bénéfiques.

Ensuite, si nous souhaitons déployer des conteneurs en production, Kubernetes est quasiment le seul orchestrateur reconnu. Il a été choisi par beaucoup d'entreprises comme Github, et ils ont commencé à créer un écosystème autour de Kubernetes. Je pense que nous devons utiliser cet écosytème et utiliser les meilleurs outils pour résoudre nos problèmes. Dans le futur nous pouvons aussi y contribuer.

## Il y a beaucoup d'outils autour de k8s, qu'est-ce que vous utilisez ?

Tout le monde a entendu parler de [Istio](https://istio.io/) mais c'est encore jeune; nous pensons l'utiliser dans un futur proche.

Il y a par exemple [kube-lego](https://github.com/jetstack/kube-lego): c'est un controleur pour [Let's Encrypt](https://letsencrypt.org/). Si tu crées un nouvel [Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/), il récupère automatiquement le certificat dont il a besoin.

Un autre bon outil est [external-dns](https://github.com/kubernetes-incubator/external-dns), c'est un autre controleur. Quand tu crées un nouvel Ingress, il récupérera automatiquement le nom de domaine auprès de [Route 53](https://aws.amazon.com/route53/) ou [Cloud DNS](https://cloud.google.com/dns/).

Enfin, il y a aussi [Heptio Ark](https://github.com/heptio/ark), un outil de reprise après sinistre. Il fait continuellement des sauvegardes du manifeste Kubernetes sur [Google Cloud Storage](https://cloud.google.com/storage/) ou [AWS S3](https://aws.amazon.com/s3/) et peut tout reconstruire lorsqu'un évènement grave arrive.

## Comment vous gérez vos bases de données ?

Nous utilisons des bases de données managées, nous n'utilisons pas Kubernetes pour ça. Kubernetes sert uniquement aux applications et services qui sont sans état (stateless).

## Comment faire pour apprendre du bon Go ?

Quand j'ai commencé à faire du Go, j'ai étudié avec le code de chez [CoreOS](https://github.com/coreos). Leur code est très bon et j'ai appris beaucoup.

Il y a aussi [HashiCorp](https://github.com/hashicorp) qui a du très bon code.

## As-tu une citation favorite ?

J'aime beaucoup la phrase de Rob Pike, «Readable means reliable» (Lisible signifie fiable). C'est la chose la plus importante quand tu écris du code.

## Quelque chose à ajouter ?

Nous recrutons au Japon, aux États-Unis et au Royaume-Uni ! Si vous êtes intéressés, [contactez-nous](https://careers.mercari.com/) !

## Merci beaucoup !

Merci à toi !
