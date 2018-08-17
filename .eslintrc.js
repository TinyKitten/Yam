module.exports = {
  root: true,
  env: {
    browser: true,
    node: true
  },
  parserOptions: {
    parser: 'babel-eslint'
  },
  extends: [
    'eslint:recommended',
    // https://github.com/vuejs/eslint-plugin-vue#priority-a-essential-error-prevention
    // より厳しいルールにするには`plugin:vue/strongly-recommended` もしくは `plugin:vue/recommended` に切り替えることを検討してください。
    'plugin:vue/recommended',
    'plugin:prettier/recommended'
  ],
  // *.vue files を lint にかけるために必要
  plugins: ['vue'],
  // ここにカスタムルールを追加します。
  rules: {
    'no-console': 'off',
    'vue/max-attributes-per-line': 'off',
    'prettier/prettier': [
      'error',
      { semi: true, trailingComma: 'es5', singleQuote: true }
    ]
  }
};
