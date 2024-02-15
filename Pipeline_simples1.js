language: node_js
node_js:
  - 14

jobs:
  include:
    - stage: test
      script: npm test

    - stage: deploy
      if: branch = main
      script:
        - npm run build
        - npm run deploy
