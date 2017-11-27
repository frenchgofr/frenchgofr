---
title: "Interview JBD dotGo 2017"
date: 2017-11-27T13:18:41+01:00
draft: false
tags : [
    "interview",
    "jbd",
    "rakyll",
]
categories : [
    "interview",
]
---

![dotgo](/img/post/171127_dotgo_jbd.jpg)

Nous avons rencontré Jaana B. Dogan, ingénieure chez Google pendant la conférence [dotGo](https://www.dotgo.eu/). Découvrons son rôle dans le développent des outils tels que pprof et sa vision du Go.

Ce billet est une traduction de l'article en anglais sur [le blog de SFEIR](https://lemag.sfeir.com/interview-jaana-b-dogan-go/).

**Bonjour JBD, première question: qui es-tu ?**

Je suis Jaana B. Dogan, développeuse logiciel chez Google. Je travaille actuellement sur notre framework d'instrumentation. C'est le framework central où nous collectons les statistiques, les traces etc. C'est aussi la couche de propagation de contexte pour tous nos langages.

**Comment c'est passée ta journée ?**

Bien, et j'étais très occupée ! Pendant que j'essayais de me préparer pendant la matinée, j'étais très anxieuse et mes slides n'étaient pas prêtes. J'avais 2 versions de la présentation: une de 10 minutes et une de 30 minutes et c'était impossible pour mois d'arriver à 18 minutes ! C'était un peu frustrant.

**Peux-tu nous expliquer ce qu'est un jour typique pour toi ?**

En quelques sortes je dirige le projet et c'est un bon mélange de négociation avec d'autres personnes, d'écriture de code et un peu d'expérimentation. Aussi, J'écris aussi des documents, je fais des revues de code, ... J'aime beaucoup ça car j'ai de la liberté et de la flexibilité pour faire plein de choses différentes. Certaines de ces taches reviennent comme des cycles et elles définissent le contenu de mes journées.

**Où en sommes-nous sur les outils ?**

Dans Go 1.10 il y aura une interface utilisateur Web pour [pprof](https://golang.org/pkg/net/http/pprof/), ce qui permettra aux utilisateurs d'avoir une vue sur pprof lorsque le flag `-http` est activé, ça fonctionnera comme l'outil lui-même. 

{{< tweet 929504504352878592 >}}

Nous avons les flame graphs, une contribution de Netflix, ce qui est génial ! Beaucoup de gens utilisent les flame graphs comme le point d'entrée pour le profiling, et c'est pour ça que c'est une bonne contribution, car elle permettra de savoir ce qui ce passe et ce sans effort. Ça abaisse la barrière d'entrée.  

{{< tweet 908748397305225217 >}}

**À propos de Go, quelle est ta fonctionnalité préférée ?**

Je pense que ce sont les fonctionnalités pour la concurrence, même si ce qui se passe en interne est parfois déroutant. Il y a des tonnes d'optimisations que j'ai découvertes quand je cherchais à en savoir plus à propos du scheduler. Il y a des tas de hooks qui optimisent les choses car le [scheduler](https://rakyll.org/scheduler/) connait tout à propos de ce qui est en train de se passer, ainsi que la concurrence, les restrictions etc. La plupart des fonctionnalités liées à la concurrence font que j'ai commencé à écrire du code en Go au départ.

**Qu'attends-tu pour Go 2 ?**

Une meilleure gestion des erreurs, nous devons faire quelque chose à ce propos. De façon générale, nous ne sommes pas efficace sur ce sujet. Beaucoup d'utilisateurs font remonter les erreurs et ne nous leur fournissons pas les primitives pour comprendre la pile d'appels actuelle, ce que peut être l'erreur et comment ils peuvent la gérer. De plus, nous ne fournissons pas toutes les informations de diagnostic nécessaires, comme la stacktrace, les pointeurs etc. Donc il y a 2 possibilités: soit gérer l'erreur, soit la remonter, et nous ne sommes efficaces dans aucun des deux. J'espère donc que nous améliorerons ça.

J'aimerai aussi beaucoup voir arriver les generics sous une forme qui reste cohérente avec le système de types, mais qui sera certainement une réécriture conséquente de l'ensemble. Cette possibilité concerne l'ensemble des utilisateurs à cause de la possibilité de fragmentation.

**Comment as-tu fini dans les parties «basses» du Go, comme le scheduler ?**

Je ne travaille pas directement sur le scheduler mais sur les outils de diagnostic, juste pour clarifier ! Mais parfois, je dois creuser pour comprendre le sens des choses et ce que nous avons besoin d'instrumenter. Précédemment j'ai travaillé dans des équipes «plateformes» et que le projet Go est devenu un succès et que j'en ai eu la possibilité, je l'ai rejoint. Je pensais au départ que mes points forts seraient le design d'API et l'outillage à cause de mon historique sur les plateformes. Je me suis passionnée pour les diagnostics très récemment et ça a contribué beaucoup à ce que je suis, je veux dire que je voulais juste sortir de ma zone de confort et apprendre.

**Est-ce que tu as une citation préférée ?**

Oui, celle de la présentation de [Sam Boyer](https://twitter.com/sdboyer) !

    Design depends largely on constraints — Charles Eames

C'est ma nouvelle citation préférée ! L'idée que les contraintes dirigent le design est une chose tellement essentielle et je pense qu'on la sous estime. Particulièrement pour les développeurs, nous aimons faire des abstractions et créer des barrières entre nous et la réalité.

**Comment tu convaincrais quelqu'un de travailler avec Go ?**

D'abord j'essaierai de comprendre quel type de problèmes ils essaient de résoudre et voir quelles parties de Go s'alignent avec la résolution. Je pense que c'est frustrant et loin de la réalité de suggérer Go à quelqu'un si les problèmes qu'ils essaient de résoudre ne sont pas en accord avec notre écosystème ou les solutions que l'on apporte.

Une fois que nous sommes d'accord sur le fait que Go peut être une bonne solution, la conversation évoluera vers «voici les choses que je souhaite faire» et «comment puis-je me former ?», alors la personne va suivre le chemin typique de l'apprentissage de Go: la documentation, le [Go Tour](https://tour.golang.org/), [Effective Go](https://golang.org/doc/effective_go.html), tout ce qui a été écrit, et sur le design d'API. Je pense que nous avons de la chance en terme de design d'API, de lisibilité, et il y a toujours des resources que je suggère aux personnes que je rencontre. Pour ma part, j'ai trouvé très utile d'avoir mon code relu, et j'ai eu beaucoup de chance car [Brad Fitzpatrick](https://twitter.com/bradfitz) m'a fait des retours sur le code que j'ai produit ! J'ai réussi à apprendre Go si bien parce que j'ai eu de la chance. Contribuer au projet d'un contributeur Go est génial aussi car tu apprends beaucoup du style de code et de la façon de penser.

**Super, merci beaucoup pour ces réponses !**

Merci à toi !

**Liens**

- [JBD on Twitter](https://twitter.com/rakyll)
- [Son blog](https://rakyll.org/)
- [À propos de sa présentation (by sourcegraph)](https://about.sourcegraph.com/go/gos-work-stealing-scheduler/)
