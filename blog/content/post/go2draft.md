---
title: "Que peut-on attendre de la version 2 de Go ?"
description: "La Gophercon 2018 qui se tenait à Denver fût l'occasion d'en savoir un peu plus sur ce qui se profile à l'horizon pour le langage."
date: 2018-08-30T09:00:00+01:00
draft: false
tags : [
    "specification",
    "annonce",
    "go2",
]
categories : [
    "article",
]
---

La [GopherCon 2018](https://www.gophercon.com/home) s'est déroulée du 27 au 30 Août 2018 à [Denver](https://www.google.fr/maps/place/Denver,+Colorado,+%C3%89tats-Unis/@39.7642548,-104.9951942,11z/data=!3m1!4b1!4m5!3m4!1s0x876b80aa231f17cf:0x118ef4f8278a36d6!8m2!3d39.7392358!4d-104.990251) dans le Colorado. 
Evénement incontournable de la communauté Go, elle nous a permis d'en savoir un peu plus sur l'avenir de notre langage favori et plus précisément d'entrevoir les nouveautés de cette mouture v2.

C'est à travers cette courte vidéo diffusée lors de la conférence, que les participants ont pu découvrir les 3 thématiques majeures de cette release v2. 

{{< youtube 6wIP3rO6On8 >}}

On y apprend que le sondage réalisé en 2017 auprès de la communauté a mis en exergue 3 attentes :

* Une meilleure gestion des **dépendances** 
* Un traitement des **erreurs** plus pratique et plus complet
* L'ajout des types **génériques**

# Dépendances

On a vu émerger [Dep](https://golang.github.io/dep/), qui s'est imposé comme le gestionnaire de dépendances de référence jusqu'à ces dernières semaines.
Finalement une autre approche à vue le jour avec l'apparition des modules.

La version de Go 1.11 récemment sortie, propose d’ores et déjà cette nouvelle fonctionnalité.
Vous trouverez plus d'informations à ce sujet dans cet [article de Dave Cheney](https://dave.cheney.net/2018/07/14/taking-go-modules-for-a-spin). 

# Gestion des erreurs

Concernant la gestion des **erreurs**, vous qui lisez ces lignes connaissez bien le pattern suivant :

{{< highlight go >}}
if err != nil {
	return err
}
{{< / highlight >}}

C'est un pattern de gestion d'erreur simple, en effet, les erreurs sont traitées comme des valeurs qu'on retourne.
Le revers de le médaille c'est qu'on retrouve ce pattern très — *voire trop* — souvent dans notre code, cela l'alourdit et pénalise la lisibilité.

Le draft de Go2 propose une approche un peu différente en introduisant un nouveau mot clé `check`.

{{< highlight go >}}
func printSum(a, b string) error {
	handle err {
		return fmt.Errorf("printSum(%q + %q): %v", a, b, err)
	}
	x := check strconv.Atoi(a)
	y := check strconv.Atoi(b)
	fmt.Println("result:", x + y)
	return nil
}
{{< / highlight >}}

On obtient un code plus **lisible** tout en permettant d'ajouter plus facilement du **contexte** à notre erreur.

Si vous voulez en savoir plus, lisez la proposition du draft sur [la gestion des erreurs](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md).

D'ailleurs en parlant de contexte, un autre axe d'amélioration des erreurs propose d'ajouter plus aisément des informations contextuelles,
on retrouve les détails sur cette proposition concernant [les valeurs des erreurs](https://go.googlesource.com/proposal/+/master/design/go2draft-error-values-overview.md).

# Types génériques

La dernière proposition, mais pas des **moindres**, traite de l'ajout des **types génériques**.
Vaste sujet s'il en est, car depuis que je fais du Go, j'ai toujours observé du coin de l'œil la communauté [débattre autour de ce sujet](https://www.reddit.com/r/golang/comments/6pgcqz/are_we_there_yet_the_go_generics_debate/). 

En effet, il est souvent reproché à Go de ne pas proposer cette fonctionnalité, même si des solutions alternatives existent aujourd'hui.
Il semble que la communauté ait été entendue ! Cette fois c'est le mot clé `contract` qui fait son apparition.
Les débats sont déjà ouverts sur leur ressemblance avec `interface`. 

{{< highlight go >}}
contract stringer(x T) {
	var s string = x.String()
}

func Stringify(type T stringer)(s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}
{{< / highlight >}}

On voit dès lors comment on pourrait désormais mutualiser du code, alors qu'aujourd'hui il nous faudrait créer une fonction pour chaque type que l'on souhaite utiliser. 

Vous retrouverez les détails de cette proposition dans cette page consacrée à [l'introduction des contrats](https://go.googlesource.com/proposal/+/master/design/go2draft-contracts.md).


Que pensez-vous de ces évolutions ? N'hésitez pas à réagir ci-dessous !



