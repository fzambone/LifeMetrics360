const { override } = require('customize-cra');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const ModuleFederationPlugin = require('webpack/lib/container/ModuleFederationPlugin');

module.exports = override(
    (config) => {
        config.plugins.push(
            new ModuleFederationPlugin({
                name: 'auth_mfe',
                filename: 'remoteEntry.js',
                exposes: {
                    './SignIn': './src/SignIn',
                    './SignUp': './src/SignUp',
                },
                shared: {
                    react: {
                        singleton: true,
                        eager: false,
                        requiredVersion: "18.2.0",
                    },
                    'react-dom': {
                        singleton: true,
                        eager: false,
                        requiredVersion: "18.2.0",
                    },
                },
            }),
            new HtmlWebpackPlugin({
                template: './public/index.html',
            }),
        );
        return config;
    }
);