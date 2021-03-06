## RAGE (React Atomic Generator)

### Purpose
My daily routine implements Atomic Pattern to the react app, I was doing that repeatedly, so the reason I create this package is to help me work every day.


### Installation
**NPM**
```bash
npm i -g @xxidbr9/rage
```
**Yarn**
```bash
yarn global add @xxidbr9/rage
```

### How To Run
```bash
rage -h
```
#### MAIN TODO
- [ ] CI/CD

#### TODO
- [ ] Init Command
- [ ] Check is react exist in package json
- [ ] Use babel / vite and inject alias
- [ ] Add alias to tsconfig.json
- [ ] Read to rage.config.[js|json] / .ragerc as a json
- [ ] Use Template folder when it exist

##### COMPONENTS
- [ ] rage create [-c | --component] {{COMPONENT_TYPE}} [--js] [--story | -s] [--name | -n | no flags] {{COMP_NAME}} 
- [X] Create Atoms Components
- [X] Create Molecules Components
- [X] Create Organisms Components
- [X] Create Templates Components

##### Misc
- [ ] Create StoryBook For Each Components when it use the flags --story, -s
- [ ] Create Test file with --test, -t flags
- [X] Show Version --version, -v
- [ ] Use Javascript instead of typescript --js


