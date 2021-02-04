1. Implémenter le mécanisme de comptage de cycles simulés
2. Combien d'instructions par seconde votre simulateur est-il capable d'exécuter ? Faites varier l'algorithme, de manière à établir quelques statistiques.
3. Soit un algorithme embarqué nécessitant un temps de traitement, sur cible, de 220 ms, sur un processeur fonctionnant à 800 Mhz. Combien de temps faut-il pour le simuler sur votre ISS ?
4. On suppose désormais que l'on cherche à étudier la possibilité d'utiliser l'ISS couplé à un systèmes réel : il s'agit alors d'uns simulation dite hybride, où certains sont réels et d'autres restent virtuels (comme notre ISS). Dans l'industrie, on parle de "hardware in the loop (HIL)" et/ou "model-in-the-loop (MIL)". On suppose que la zone mémoire de données (buffer) est écrite par un démultiplexeur satellitaire, qui envoit un flux continu d'images numériques, au rythme de n images HD par secondes. Chaque image doit être traitée dans le temps de transmission d'une image : ceci signifie que pendant que le "demux" transmet 1 image, la précédente est en cours de traitement et qu'il n'y a aucune latitude à empiéter sur le temps alloué à chaque image. Ceci serait possible avec une FIFO, mais on se l'interdit ici. Ce traitement "tendu" se réalise avec un simple buffer ping-pong, qui laisse très peu de flexibilité en matière algorithmique. Question : quel est le nombre m atteignable de cycles de traitement sur ISS dans cette simulation hybride ? Quel pourcentage de l'image serait alors traité par cet ISS ? Conclure. [cet exercice nécessite de rechercher sur Internet certaines grandeurs.]

// CACHES : https://moodle.ensta-bretagne.fr/course/view.php?id=1292


## A implémenter

* cache
* interruption