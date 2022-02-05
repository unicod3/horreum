# ![Logo](logo.png) Horreum

From Wikipedia, the free encyclopedia;

> *A **horreum** (plural: horrea) was a type of public warehouse used during the ancient Roman period.*
> Although the Latin term is often used to refer to granaries, Roman horrea were used to store many other types of consumables; the giant Horrea Galbae in Rome were used not only to store grain but also olive oil, wine, foodstuffs, clothing and even marble...
> 
> https://en.wikipedia.org/wiki/Horreum

Following the ancient Roman correspondence Horreum, is an application to manage products and their stock informations.


## Architectural Design

Every application prone to the architectural changes while time goes on and needs changes
and Horreum no exception to that, to address those concerns Horreum uses [C4 Modelling](https://c4model.com/).

In order to keep the architectural changes manageable over time, I've used [C4-PlantUML](https://github.com/plantuml-stdlib/C4-PlantUML/)
to document architectural structure, they can be found under `docs/c4model` folder.

For the sake of the demo I only designed `System Context` and `Container` Level.

| 1.System Context                                            |                            2.Horreum Container                             |
|-------------------------------------------------------------|:--------------------------------------------------------------------------:|
| <img src="./docs/c4model/SystemContext.png" alt="drawing"/> | <img src="./docs/c4model/HorreumContainer.png" alt="drawing" width="550"/> |

 
## Project structure

Horreum uses [Standard Go Project Layout](https://github.com/golang-standards/project-layout) for the structural design.



