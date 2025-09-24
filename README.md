## Perfect21 Documentation

### Overview

Perfect21 is a web application designed to help users memorize the perfect strategy for H17 Blackjack. The app determines the optimal action based on player and dealer cards, utilizing a JSON-based strategy chart and a Go-based API.

 

**Main Features** [0:15](https://loom.com/share/e1b69c6329914a20b0c9aa5827a63c8a?t=15)

![generated-image-at-00:00:15](https://loom.com/i/1062842f2a554893b1a7e9acb189b98d?workflows_screenshot=true)

- Memorizes the perfect strategy for H17 Blackjack.
- Determines the optimal action based on player total and dealer upcard.
- Provides feedback on user actions (correct or incorrect).

 

**Gameplay Example** [0:44](https://loom.com/share/e1b69c6329914a20b0c9aa5827a63c8a?t=44)

![generated-image-at-00:00:44](https://loom.com/i/1dc388227d87491b818129efabb98026?workflows_screenshot=true)

- Users can input their cards and the dealer's upcard.
- The app evaluates the player's action against the optimal strategy.
- Example scenarios include: 
  - Player has Queen + 3 against dealer's 5 + hidden card.
  - Correct actions include Hit, Stand, or None based on the strategy chart.

 

**API Overview** [2:49](https://loom.com/share/e1b69c6329914a20b0c9aa5827a63c8a?t=169)

![generated-image-at-00:02:49](https://loom.com/i/ddfa2c3d6f194b7eb4a3db5c50b5f0d8?workflows_screenshot=true)

- The backend API is built using Go with the Gin framework.
- It has a single GET route: `/play`.
- The API returns a JSON object containing: 
  - Player cards
  - Dealer cards
  - Correct action based on the strategy chart.

 

**Frontend Implementation** [5:36](https://loom.com/share/e1b69c6329914a20b0c9aa5827a63c8a?t=336)

![generated-image-at-00:05:36](https://loom.com/i/56ebc3955f7a40c8a6d4484de9f59cd5?workflows_screenshot=true)

- The frontend is developed using Vite and React.
- Utilizes state management with hooks (useState) to manage game state.
- Displays results and correct/incorrect feedback based on user actions.

 

**Future Enhancements** [7:00](https://loom.com/share/e1b69c6329914a20b0c9aa5827a63c8a?t=420)

![generated-image-at-00:07:00](https://loom.com/i/70853944e62a49749c7850d61a4cc19d?workflows_screenshot=true)

Potential features for future development include:

- Streak tracking for correct answers.
- Full Blackjack gameplay with win/loss outcomes.
- Bankroll management features.
- Options to toggle strategy tips and rule sets for different casinos.

### Link to Loom

<https://loom.com/share/e1b69c6329914a20b0c9aa5827a63c8a>

# React + TypeScript + Vite

This template provides a minimal setup to get React working in Vite with HMR and some ESLint rules.

Currently, two official plugins are available:

- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react) uses [Babel](https://babeljs.io/) for Fast Refresh
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react-swc) uses [SWC](https://swc.rs/) for Fast Refresh

## React Compiler

The React Compiler is enabled on this template. See [this documentation](https://react.dev/learn/react-compiler) for more information.

Note: This will impact Vite dev & build performances.

## Expanding the ESLint configuration

If you are developing a production application, we recommend updating the configuration to enable type-aware lint rules:

```js
export default defineConfig([
  globalIgnores(['dist']),
  {
    files: ['**/*.{ts,tsx}'],
    extends: [
      // Other configs...

      // Remove tseslint.configs.recommended and replace with this
      tseslint.configs.recommendedTypeChecked,
      // Alternatively, use this for stricter rules
      tseslint.configs.strictTypeChecked,
      // Optionally, add this for stylistic rules
      tseslint.configs.stylisticTypeChecked,

      // Other configs...
    ],
    languageOptions: {
      parserOptions: {
        project: ['./tsconfig.node.json', './tsconfig.app.json'],
        tsconfigRootDir: import.meta.dirname,
      },
      // other options...
    },
  },
])
```

You can also install [eslint-plugin-react-x](https://github.com/Rel1cx/eslint-react/tree/main/packages/plugins/eslint-plugin-react-x) and [eslint-plugin-react-dom](https://github.com/Rel1cx/eslint-react/tree/main/packages/plugins/eslint-plugin-react-dom) for React-specific lint rules:

```js
// eslint.config.js
import reactX from 'eslint-plugin-react-x'
import reactDom from 'eslint-plugin-react-dom'

export default defineConfig([
  globalIgnores(['dist']),
  {
    files: ['**/*.{ts,tsx}'],
    extends: [
      // Other configs...
      // Enable lint rules for React
      reactX.configs['recommended-typescript'],
      // Enable lint rules for React DOM
      reactDom.configs.recommended,
    ],
    languageOptions: {
      parserOptions: {
        project: ['./tsconfig.node.json', './tsconfig.app.json'],
        tsconfigRootDir: import.meta.dirname,
      },
      // other options...
    },
  },
])
```
