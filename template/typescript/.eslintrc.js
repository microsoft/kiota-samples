module.exports = {
  root: true,
  env: {
    browser: true,
    es6: true,
    node: true,
  },
  parser: "@typescript-eslint/parser",
  parserOptions: {
    project: "tsconfig.json",
    sourceType: "module",
  },
  plugins: ["@typescript-eslint"],
  rules: {
    "@typescript-eslint/quotes": "off",
    "@typescript-eslint/indent": "off",
    "@typescript-eslint/semi": "off",
    "@typescript-eslint/member-delimiter-style": [
      "error",
      {
        multiline: {
          delimiter: "comma",
          requireLast: true,
        },
        singleline: {
          delimiter: "comma",
          requireLast: false,
        },
      },
    ],
    "comma-dangle": ["error", "always-multiline"],
    "max-classes-per-file": "off",
    "no-console": "error",
    "no-multiple-empty-lines": ["error", { max: 1 }],
    "no-redeclare": "error",
    "no-return-await": "off",
    "prefer-const": "error",
  },
};
