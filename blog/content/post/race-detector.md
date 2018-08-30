---
title: "Le détecteur de situation de compétition (race condition detector) de Go trouve-t-il tous les conflits?"
description: "Une traduction d'un article de Valentin Deleplace (Google)"
date: 2018-08-29
draft: false
tags : [
    "article",
    "Golang",
    "Programming",
]
categories : [
    "article",
]
---

Ceci est une traduction du billet sur [Medium de Valentin Deleplace](https://medium.com/@val_deleplace/does-the-race-detector-catch-all-data-races-1afed51d57fb).

En bref: il détecte toutes les situations de compétition au moment où elles surviennent.

Le [détecteur de situation de compétition](https://golang.org/doc/articles/race_detector.html) (Data Race Detector) est une des fonctionnalités majeures de l'outillage Go.
Aujourd'hui, je souhaite clarifier la portée de son pouvoir de détection.

![](/img/post/180829_race_detector/gopher_race.gif)

## Une situation de compétition, c'est quoi ?

Une telle situation survient quand deux goroutines accèdent à la même variable de façon concurrente et au moins un de ces accès est une écriture.

## De telles situations existent en Go

Il est possible d'écrire du code incorrectement synchronisé et de le compiler. Le compilateur le considère comme valide et n'émet pas d'avertissements ou d'erreurs concernant la synchronisation.

Cependant, ces bogues sont réels et entraînent des comportements inattendus, qui peuvent être un plantage, un comportement incorrect et silencieux ou un comportement correct et silencieux «par hasard».
Voici [quelques exemples](https://golang.org/doc/articles/race_detector.html#Typical_Data_Races).

Il est fondamental de comprendre qu'une ***situation de compétition*** est une propriété d'une exécution particulière et qu'un ***bogue de synchronisation*** est une propriété du programme.

Une situation de compétition est toujours due à un bogue, mais un bogue peut rester silencieux (ne pas entraîner de situation de compétition) lors de multiples exécutions.

## Le détecteur n'attrape pas les bugs au moment de la compilation

```bash
$ go build -race mycmd
```

Cette commande ***ne signifie pas*** «Compile `mycmd`, et donne-moi les bogues de synchronisation que tu as trouvés pendant la compilation».

Elle signifie «Compile une version spéciale et instrumentée de `mycmd`. Quand `mycmd` sera lancé, si une situation de compétition a lieu pendant l'exécution,
affiche un avertissement sur la sortie standard (stdout) et à la fin sort du programme avec un code 66».

Cette version instumentée a [un coût à l'exécution](https://golang.org/doc/articles/race_detector.html#Runtime_Overheads), autrement dit elle est plus lente et utilise plus de mémoire.

## Le détecteur m'a affiché un avertissement mais je ne pense pas que mon code comporte un bogue. Je peux l'ignorer ?

Non.

Le détecteur ne produit pas de faux positifs.

Quand il émet un avertissement, cela signifie toujours qu'une situation de compétition a eu lieu.

Quand une telle situation a lieu, cela signifie toujours que le programme comporte un bogue.

Si vous pensez réellement que vous avez trouvé un faux positif, alors remontez le bogue pour le détecteur de situation de compétition.
Si vous avez de bonnes raisons de croire que la situation a pour cause la librairie standard ou l'environnement d'exécution (plutôt que votre code), remontez le bogue au niveau de ceux-ci.

## Le détecteur a affiché un avertissement venant d'une partie non critique de mon code. Je peux l'ignorer ?

Non.

Lire [Benign Data Races: What Could Possibly Go Wrong?](https://software.intel.com/en-us/blogs/2013/01/06/benign-data-races-what-could-possibly-go-wrong).

![Le gaufre aviateur pense qu'il peut soudainement occuper la même position que l'arbre, alors qu'en fait c'est problématique.](/img/post/180829_race_detector/gopher_aviator.png)

> Le [gaufre](https://fr.m.wikipedia.org/wiki/Gaufre_(animal)) aviateur pense qu'il peut soudainement occuper la même position que l'arbre, alors qu'en fait c'est problématique.

## Je lance mon programme et le détecteur ne me remonte rien. Cela signifie-t-il que mon code est correctement synchronisé ?

Non.

Cela signifie que pendant l'exécution du programme, la séquence particulière de lectures et écritures entrelacées effectivement produites n'a pas violé le [modèle de mémoire](https://golang.org/ref/mem).
Ça ne sous-entend pas qu'une autre exécution du même programme ne créera pas une situation de compétition.

De nombreux facteurs peuvent faire que la trace d'exécution peut différer entre plusieurs exécutions.

Dans le programme suivant, nous avons 3 compteurs.

![](/img/post/180829_race_detector/counter.png)

2 goroutines incrémentent chacune un compteur, et la goroutine principale attend qu'elles aient terminé. Mais les 2 goroutines peuvent choisir d'incrémenter le même compteur !

{{< gist Deleplace 0bfbf70427c1ea84cff5dc8742ca1ae1 >}}

Considérez maintenant ces 2 exécutions:

```bash
$ go run -race racy3.go 
[1 0 1]
```

Et ensuite:

```bash
$ go run -race racy3.go 
==================
WARNING: DATA RACE
Read at 0x00c4200b8000 by goroutine 7:
  main.main.func2()
      /home/deleplace/racy3.go:30 +0x66
Previous write at 0x00c4200b8000 by goroutine 6:
  main.main.func1()
      /home/deleplace/racy3.go:24 +0x87
Goroutine 7 (running) created at:
  main.main()
      /home/deleplace/racy3.go:28 +0x1d0
Goroutine 6 (finished) created at:
  main.main()
      /home/deleplace/racy3.go:22 +0x192
==================
[0 0 2]
Found 1 data race(s)
```

Lors de la première exécution, les nombres aléatoires ont eu pour valeurs 0 et 2, et donc les goroutines ont incrémenté des compteurs différents.

![](/img/post/180829_race_detector/gopher_race_ok.png)

Dans la seconde exécution, les nombres aléatoires ont eu la même valeur 2, et donc les goroutines ont incrémenté le même compteur sans mécanisme de synchronisation adapté (pas de relation [happens-before](https://golang.org/ref/mem#tmp_2)),
ce qui est interdit et est un problème grave.

![](/img/post/180829_race_detector/gopher_race_concurrent.png)

Détecter les bogues de synchronisation pour toutes les exécutions possibles d'un programme Go solutionnerait [le problème de l'arrêt](https://fr.wikipedia.org/wiki/Probl%C3%A8me_de_l%27arr%C3%AAt),
ce qui est hors du périmètre du détecteur de situation de compétition.

## Si une situation de compétition se produit pendant l'exécution de mon programme, le détecteur la trouvera-t-elle ?

Oui.

En instrumentant le fichier binaire généré pour vérifier chaque écriture et chaque lecture dans la mémoire, le détecteur réagit à tous les accès mal synchronisés, au moment où ils ont lieu.

## Conclusion

- Le détecteur de situation de compétition est simple à utiliser. Je le recommande chaudement.
- Il ne peut pas trouver tous les bogues qui se cachent dans votre code.
- Mais il lance des **AVERTISSEMENTS** pour chaque situation de compétition qu'il rencontre.
- Toutes les situations: pas de faux négatifs !
- Seulement les vraies situations: pas de faux positifs !

![](/img/post/180829_race_detector/gopher_wow.png)

Merci à [François Beaufort](https://medium.com/@beaufortfrancois?source=post_page) et [Laurent Picard](https://medium.com/@PicardParis?source=post_page).
