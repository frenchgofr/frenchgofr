---
title: "Quel est le mot-clé le plus courant dans la librairie standard Go ?"
description: "Une traduction d'un article de Francesc (Just for func)"
date: 2017-12-05T06:49:41+01:00
draft: false
tags : [
    "article",
    "francesc",
    "Golang",
    "Scanner",
    "justforfunc",
    "Programming",
]
categories : [
    "article",
]
---

Ceci est une traduction du billet sur [Medium de Francesc](https://medium.com/@francesc/whats-the-most-common-identifier-in-go-s-stdlib-e468f3c9c7d9).

Ce billet est la transcription du dernier épisode de [justforfunc](http://justforfunc.com/) titré «What’s the most common identifier in Go’s stdlib?». Le code pour le programme se trouve [ici](https://github.com/campoy/justforfunc/blob/master/24-ast/scanner/main.go), dans le dépôt de [justforfunc](https://github.com/campoy/justforfunc).

{{< youtube k23xhJoTbI4 >}}

Si vous préférez regarder une vidéo à lire un article ;)

## Le problème

Imaginez avoir le programme suivant et que vous souhaitiez en extraire tous les mots-clés.


{{< gist campoy 1eda05a8db6a5f97d7effbf9490e9638 >}}

Nous devons obtenir une liste contenant `main`, `fmt` et `Println`.

## Qu'est-ce qu'un mot-clé ?

Pour répondre à cette question il faut plonger un peu dans la théorie des langages. Juste un peu, pas d'inquiétude.

Les langages de programmation sont entre autres définis pas une série de règles qui définissent un programme valide. Ces règles ressemblent à ceci:

{{< highlight lexx >}}
IfStmt = "if" [ SimpleStmt ";" ] Expression Block [ "else" ( 
IfStmt | Block ) ] .
{{< /highlight >}}

Cette règle nous indique ce à quoi l'instruction `if` ressemble en Go. Les parties `"if"`, `";"`, et `"else"` sont les mots-clés de la règles qui nous aide à comprendre la structure du programme, alors que `ExpressionBlock`, `SimpleStmt`, etc sont d'autres règles.

L'ensemble de ces règles est appelé la grammaire du langage. Vous pouvez trouver l'ensemble dans la [spécification du langage Go](https://golang.org/ref/spec).

Ces règles ne sont pas définies pour les caractères présents dans le programme, mais sur des lexèmes (tokens en anglais). Ceux-ci sont atomiques comme `if` ou `else`, mais peuvent aussi être de type complexe comme les nombres entiers `42`, les nombres flottants `4.2`, les chaînes de caractères `"hello"`, ou... les mots-clés tels que `main`.

Comment sait-on que `main` est un mot-clé et non un nombre ? Et bien il y a aussi des règles pour ça. En lisant la partie [identifiers](https://golang.org/ref/spec#identifier) de la spécification du langage, vous trouverez cette règle:

{{< highlight lexx >}}
identifier = letter { letter | unicode_digit } .
{{< /highlight >}}

Dans cette règle, `letter` et `unicode_digit` ne représentent pas des lexèmes; ils sont une classe de caractères. Donc étant donné l'ensemble de ces règles, il devient presque simple d'écrire un programme qui va lire caractère par caractère et qui «émet» un lexème lorsqu'il détecte un groupe de caractères qui correspond à une règle.

Si nous commençons avec : `fmt.Println` ça doit générer les lexèmes : `fmt` en tant que mot-clé, `.`, et `Println` en mot-clé. Est-ce un appel de fonction ? À ce point nous ne pouvons pas le savoir, et ce n'est pas important. La seule structure est la séquence qui nous laisse découvrir l'ordre dans lequel les choses apparaissent.

![Étant donné une séquence de caractères, un scanner génère une séquence de lexèmes](/img/post/mots-cles_lexemes.png)

Ce type de programme qui donne une séquence de lexèmes à partir d'une séquence de caractères est appelé un scanner. La librairie standard de Go fournit un scanner pour les programmes Go dans [`go/scanner`](https://golang.org/pkg/go/scanner/). Les types de lexèmes qu'il génère sont définis dans [`go/token`](https://golang.org/pkg/go/token/).

## Utiliser go/scanner

Ok, on sait ce qu'est un scanner. Comment l'utilise-t'on ?

### Lire les arguments depuis la ligne de commande

Commençons par un petit programme qui affiche simplement tous les arguments passés lorsqu'on l'exécute. Ce sera notre point de départ.

{{< gist campoy ac42675c871c7587578f409c4ca13269 >}}

(Afficher tous les arguments de la ligne de commande)

Ensuite, nous devons scanner tous les fichiers passés en argument. Pour cela nous devons créer un nouveau `scanner.Scanner` et l'initialiser avec le contenu du fichier.

### Afficher chaque lexème

Avant de pouvoir appeler la méthode `Init` du `scanner.Scanner` nous allons lire le contenu du fichier et créer un `token.FileSet` avec un `token.File` par fichier que nous scannons.

Une fois le scanner initialisé nous pouvons appeler `Scan` et afficher le lexème obtenu. Quand la fin du fichier scanné est atteinte, nous obtiendrons un lexème de fin de fichier `EOF` (End Of File).

{{< gist campoy 338870a69acc90fe48a33a519fecbfa6 >}}

### Compter les lexèmes

Super, nous pouvons afficher tous les tokens, mais nous avons besoin de garder une trace afin de savoir combien de fois nous trouvons chacun d'eux, pour les trier par occurence et afficher les 5 plus courants.

En Go, la meilleure façon de faire est d'utiliser une map dont la clé est le mot-clé et la valeur est le nombre d'occurences.

À chaque fois que l'on rencontre un mot-clé, nous devons incrémenter son compteur. À la fin, nous convertissons la map en une [slice](https://blog.golang.org/go-slices-usage-and-internals), que nous pouvons trier et afficher.

{{< gist campoy d4a147ab7d8801915033840e0e11f786 >}}

Certaines parties du code ont été retirées pour plus de clarté, le code complet se trouve [ici](https://github.com/campoy/justforfunc/blob/master/24-ast/scanner/main.go), dans le [dépôt justforfunc](https://github.com/campoy/justforfunc)

## Alors, quel est le mot-clé le plus courant ?

Lançons le programme en utilisans les contenus de `github.com/golang/go`.

{{< highlight bash >}}
$ go install github.com/campoy/justforfunc/24-ast/scanner
$ scanner ~/go/src/**/*.go
 82163 v
 46584 err
 44681 Args
 43371 t
 37717 x
{{< /highlight >}}

Le mot-clé le plus courant est `v`, tu parles de mots-clés courts ! Nous allons compter uniquement les mots-clés ayant au moins trois caractères, en modifiant légèrement le code:

{{< highlight go >}}
for s, n := range counts {
	if len(s) >= 3 {
		pairs = append(pairs, pair{s, n})
	}
}
{{< /highlight >}}

Et relançons notre programme:

{{< highlight bash >}}
$ go install github.com/campoy/justforfunc/24-ast/scanner
$ scanner ~/go/src/**/*.go
 46584 err
 44681 Args
 36738 nil
 25761 true
 21723 AddArg
{{< /highlight >}}

Rien de surprenant ici, `err` et `nil` sont présents dans quasiment tous les programmes qui contient `if err != nil`. Et pour `Args` ?

Ce sera le sujet d'un futur épisode.

## Remerciements

Si vous avez apprécié cet épisode, n'oubliez pas de souscrire à [justforfunc](http://justforfunc.com/) ! Aussi, vous pouvez soutenir la série sur [patreon](https://patreon.com/justforfunc).

Merci à [Francesc](https://twitter.com/francesc) de nous avoir autorisé à traduire son article, sa chaîne vient de passer les [10 000 abonnés](https://www.youtube.com/watch?v=kuted8F2KJw), bravo à lui !

