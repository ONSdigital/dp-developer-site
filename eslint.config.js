const babelParser = require("@babel/eslint-parser");
const pluginJest = require("eslint-plugin-jest");

module.exports = [
    {
        ignores: ["node_modules/**", "assets/**"]
    },
    {
        files: ["**/*.js"],
        languageOptions: {
            parser: babelParser,
            parserOptions: {
                requireConfigFile: false,
                babelOptions: { babelrc: false, configFile: false }
            },
            globals: {
                window: true, document: true,
                require: true, module: true,
            }
        },
        plugins: { jest: pluginJest },
        rules: {
            ...pluginJest.configs.recommended.rules,
            "max-len": [2, { code: 140, tabWidth: 4, ignoreUrls: true }],
            "require-jsdoc": 0,
        }
    }
];