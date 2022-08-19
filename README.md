# Edge Computing

The project aims to illustrate [the edge computing](https://en.wikipedia.org/wiki/Edge_computing). It provides the web application executing entire business logic on the client side to eliminate the client-server communication as the crucial dependency of the UX flow.

## Demo Application

The application is a spin-off of [the color-theory-app](https://github.com/kislerdm/color_theory_app). Its logic returns the color's _name_ and [_type_](https://en.wikipedia.org/wiki/Color_theory#Warm_vs._cool_colors) based on the user's input.

### Features breakdown

#### Color Name

[The lookup table](./data/colorsname.csv) is used to define the color name based on its RGB code:
1. [Euclidean distance](https://en.wikipedia.org/wiki/Euclidean_distance) between the input color code and the lookup table's codes is calculated in the RGB space.
2. The "name" value is returned from the table's row with the RGB point least distant from the input RGB point.

_Limitations_:

- Empty string is returned, if no row with the distance below `50.` units is found;
- First found color from the table is returned if several colors are equidistant from the input in the RGB space.

#### Color Type

The color types include two options: "_Warm_" and "_Cold_". An [xgboost](https://xgboost.readthedocs.io/en/release_1.3.0/) binary classification [model was developed](./colortypemodel/main.py) using [the data sample](./data/colortype_train.csv) generated manually (thanks to my wife for that :) ).

The trained model is [saved as JSON file](https://xgboost.readthedocs.io/en/release_1.3.0/python/python_api.html?highlight=dump_model#xgboost.Booster.dump_model).

## Means of delivery

The logic will be delivered using different technologies to assess delivery process in terms of the following metrics:

- total assets size
- application performance
- development complexity

### Codebase tree

```bash
.
|-- README.md
|-- Makefile                  <- commands to test/build/deploy all implementations
|-- public
|   |-- index.html            <- landing page when hitting https://edge-computing-demo.dkisler.com 
|   |-- go
|   |    |-- index.html       <- copied from app/common; loaded on request to https://edge-computing-demo.dkisler.com/go/
|   |    `-- assets
|   |        |-- logic.wasm   <- logic build artefact
|   |        |-- index.js     <- concatination of app/go/assets/index.js and app/common/index.js
|   |        |-- styles.css   <- copied from app/common
|   |        `-- favicon.jpg  <- copied from app/common
|   |-- js
|   |-- rust
|   |-- technologyFoo
|   |-- ...
|   `-- technologyBar
`-- app
    |-- common
    |   |-- index.html
    |   `-- assets
    |       |-- index.js
    |       |-- styles.css
    |       `-- favicon.jpg
    |-- go
    |   |-- README.md     <- readme with details about implementation
    |   |-- Makefile      <- commands to test and build the logic
    |   |-- ...           <- logic implementation
    |   `-- assets        <- build results
    |       |-- index.js  <- all js dependecies
    |       `-- ...       <- all other artefacts to be dispatched
    |-- js
    |-- rust  
    |-- technologyFoo  
    |-- ...     
    `-- technologyBar  
```

#### Requirements

`Makefile` has to include the following targets:

- test
- build

### Commands

Run to build the app and all its dependencies for deployment:

```bash
make build TECHNOLOGY=##language##
```

For example, run to build the app in Go:
```bash
make build TECHNOLOGY=go
```

#### Build flow

1. The logic is compiled and the resulting artefacts are stored to the `app/{{technology}}/assets` subdirectory. _Example_: for Go, binary `logic.wasm` is stored to `app/go/assets/logic.wasm`. 
2. All `js` dependencies are combined to `app/{{technology}}/assets/index.js`. _Example_: for Go, [`wasm_exec.js`](https://tinygo.org/docs/guides/webassembly/#how-it-works) and [`the init script`](app/go/assets/init.js) shall be combined to `app/go/assets/index.js`.
3. Common dependencies are copied from `app/common` to `public/{{technology}}`.
4. Technology-specific artefacts are copied from `app/{{technology}}/assets/` to `public/{{technology}}/assets/`.
5. `public` pushed to github artefacts as part of [CI pipeline](https://github.com/peaceiris/actions-gh-pages) to deploy static website as github-pages deployment.
