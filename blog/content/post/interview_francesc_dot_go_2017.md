---
title: "Interview Francesc dotGo 2017"
date: 2017-11-07T13:18:41+01:00
draft: false
tags : [
    "interview",
    "francesc",
]
categories : [
    "interview",
]
---

![dotgo](/img/post/171106_dotgo_francesc_ml.jpg)

On ne présente plus [francesc](https://twitter.com/francesc), le gopher le plus médiatisé de l'écosystème [Go](https://golang.org).
Nous avons eu la chance de l'interviewer lors de la dernière [dotGo](https://www.dotgo.eu/) qui s'est tenue à Paris lundi 06 novembre dernier.
Et voici les questions que nous lui avons posé.  

**Bonjour Francesc, première question, interview en Français ou en anglais ?**

C'est pour un site français ?  Alors en français.

**On entend beaucoup parler de toi dans le monde du Go, peux-tu te présenter ?**

Je suis Francesc Campoy, je suis de Barcelone et j'ai habité en France pendant 7 ans et c'est pour ça que je parle français.
J'ai travaillé chez Google pendant 7 ans, dont 5 dans l'équipe Go. Je viens de partir et j'ai rejoint une boîte qui est à Madrid, mais je suis toujours basé à San Francisco.
Je vais maintenant faire du Machine Learning sur du code source, l'idée c'est de faire du Big Data mais avec du code source comme données d'entrée.
Ils font beaucoup de Go et d'[Open Source](https://github.com/src-d), je les rejoins en tant que VP of Developer Relations.
Sinon je fais aussi des vidéos Youtube, ma chaîne s'appelle [JustForFunc](http://youtube.com/c/justforfunc) et apparamment il y a beaucoup de gens qui regardent ça, c'est génial (toute la dotGo à levé la main à cette question N.D.L.R.).
J'ai vu ça aujourd'hui et c'est vraiment génial.

**Comment es-tu venu au Go ?**

Ça c'est intéressant comme question. En fait je suis devenu "expert" C++ chez Amadeus à Sophia Antipolis et j'en ai fait pendant très très très longtemps.
Quand j'ai rejoint Google en tant que software engineer C++, j'en avais marre et je voulais faire quelque chose de nouveau.
J'ai eu l'opportunité de devenir Developer Relations et c'est quelque chose que j'ai beaucoup aimé.
A l'époque, mon sujet de master c'était un compilateur et je me suis dit que ça serait sympa de travailler en langage, ça serait un challenge intéressant.
Et comme le seul langage que Google avait c'était Go, donc j'ai commencé à l'apprendre.
Je parlais avec l'équipe Go et je leur disais que j'aimerais bien les rejoindre.
Après quelques interview avec plein de monde dont une avec [Rob Pike](https://fr.wikipedia.org/wiki/Rob_Pike) (un des créateur du langage N.D.L.R.), ils ont accepté ma candidature et j'ai pu rejoindre l'équipe Go.
Mon premier boulot dans l'équipe à été d'apprendre le Go et voir quels étaient les problèmes. J'ai essayé de les résoudre, ajouté de la documentation, des tutoriaux et plein de choses comme ça.

**La langage évolue, on [parle d'une version 2.0](https://www.youtube.com/watch?v=0Zbh_vmAKvk), quelles seraient tes envies pour une future version du langage ?**

Donc forcément avec la version 2.0 du langage, la réponse que tout le monde attends c'est : les generics. Je trouve que les generics ça serait bien mais pas forcément le plus important pour moi.
Ce que j'aimerai bien voir maintenant qu'on a [dep](https://github.com/golang/dep), un outil qui nous permet d'avoir un meilleure gestion des bibliothèques autres que celles de la bibliothèque standard,
c'est avoir une bibliothèque standard du langage qui soit plus petite.
Il y a beaucoup de choses dans la bibliothèque standard qui ne devraient pas être là, comme par exemple l'encodeur PNG ou JPG.
On en a pas vraiment besoin, c'est bien qu'elles soient là, mais s'il y a d'autres personnes qui peuvent faire mieux et plus rapide, pourquoi on ne leur laisserait pas faire ?
On donne actuellement un avantage aux API de la bibliothèque standard, ce qui n'est pas forcément pertinent et comme [contribuer au langage](https://www.youtube.com/watch?v=DjZMKKfNVMc) lui même c'est plus difficile,
ça implique une évolution plus lente de ces bibliothèques.
J'aimerais bien corriger ce point mais on verra ce que l'avenir nous réserve.

**Tu le disais précédemment, on te connait aussi pour [JustForFunc](http://youtube.com/c/justforfunc), est-ce que tu es beaucoup sollicité pour tes épisodes par les gophers ?**

Oui, j'ai un [formulaire](http://form.justforfunc.com) et je reçois beaucoup d'idées. Je fais surtout des codes reviews et les gens aiment bien ça, ils veulent que je fasse la revue de leur code.
Je dois avouer que c'est pas toujours faisable parce qu'ils envoient parfois des vrais paquet de plusieurs dizaines de milliers de lignes de code.
Mais il m'est arrivé par exemple que parmi quelque chose d'assez grand, il y avait une partie ou un package intéressant et du coup je fais la revue d'un seul paquet et c'était très intéressant.
J'ai eu la chance de travailler chez Google et particulièrement dans l'équipe Go et ça m'a permis d'avoir des revues de code qui étaient vraiment très enrichissantes.
Elles me venaient de gens comme Rob Pike et j'aimerais pouvoir renvoyer l'ascenseur à la communauté et partager cette expérience.
La façon la plus simple que j'ai trouvé de faire ça, c'est les vidéos. J'ai essayé d'écrire des articles de blog pour des revues de code et... ça marche pas (rires).
C'est beaucoup plus simple dans une vidéo de montrer comment tu changes le code.

**Quel est le niveau de préparation pour tes épisodes ? Préparés au millimètre, un peu d'improvisation quand même ?**

Il y a carrément de l'improvisation, il y a beaucoup de gens qui disent que c'est très bien que j'essaye de mettre des erreurs, je les mets pas, je les fais (rires).
Je fais des erreurs pendant que je code, c'est normal, tout le monde en fait, mais je ne les enlève pas et c'est ça aussi que les gens aiment bien.
Ca permet de montrer comment je fais quand quelque chose ne marche pas, quelle est ma façon de corriger les problèmes.
Sinon ça dépend beaucoup des épisodes. Il y a des épisodes qui m'ont demandés beaucoup de temps, par exemple un des épisodes le plus vus est celui sur le package [context](https://www.youtube.com/watch?v=LSzR0VEraWw), celui-là m'a pris plusieurs mois.
Mais parce qu'à la base c'était un talk que j'ai fait quatre fois et après j'en ai fait un épisode.
Donc ce qui m'arrive des fois c'est que je fais des talks et au lieu de les faire une nième fois, j'en fais une vidéo et ça m'aide à passer à autre chose.
Si des personnes ont des questions sur le sujet, je peux leur dire d'aller voir la vidéo.
Mon dernier épisode a été fait comme ça. J'ai fait trois talks sur l'[execution tracer](https://github.com/golang/go/issues/16526), la quatrième fois plutôt que de le faire encore sur un meetup,
j'ai produit un épisode de [Justforfunc](https://www.youtube.com/watch?v=ySy3sR1LFCQ) sur le sujet.

**Comme tu l'as annoncé tu vas chez [source{d}](https://blog.sourced.tech), est-ce que tu vas poursuivre ton travail pour la communauté Go ?**

Oui tout à fait ! Je vais même essayer de faire plus, car jusqu'à présent j'étais dans l'équipe Go mais j'étais aussi dans l'équipe [Google Cloud Platform](https://cloud.google.com),
donc je faisais beaucoup de choses qui n'étaient pas forcément liées au Go mais il fallait les faire.
Maintenant la société que je rejoins, même si elle est focalisée sur le Machine Learning, on y fait aussi énormément de Go.
Tout le backend applicatif est écrit en Go, donc mon idée c'est de continuer à apprendre le Go forcément, mais aussi d'avoir des ingénieurs qui travaillent sur le backend pour faire plus de talks.
Mon but c'est d'essayer de donner à la communauté le plus possible et peut-être que je ferai des choses un peu différentes, plus de Machine Learning,
le talk que j'ai donné aujourd'hui par exemple est plus dans l'esprit de ce que j'essaierai de faire dorénavant.
Est-ce qu'on peut utiliser les GPU avec Go ? Pour le moment j'ai pas vu grand monde qui le faisait, mais c'est intéressant, donc peut-être que je ferai ça.

**Pour finir, est-ce que tu as une punchline ou une citation préférée ?**

Intéressant mais c'est une question difficile, laisse moi réfléchir.
La réponse facile serait celle de [Rob Pike](https://go-proverbs.github.io) : _"Don't communicate by sharing memory; share memory by communicating"_, mais c'est trop simple.
Ah oui : _**"Ne pense pas à ce que la communauté Go peut faire pour toi, pense à ce que tu peux faire pour la communauté".**_

**Merci Francesc, on peut faire un selfie ?**

Oui bien sûr !

![dotgo](/img/post/171106_interview_francesc.jpg)
