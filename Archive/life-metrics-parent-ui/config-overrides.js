const { override } = require('customize-cra');
const ModuleFederationPlugin = require('webpack/lib/container/ModuleFederationPlugin');

module.exports = override(
    (config, env) => {
        config.plugins.push(
            new ModuleFederationPlugin({
                name: 'life_metrics_parent_ui',
                remotes: {
                    auth_mfe: 'auth_mfe@http://localhost:3001/remoteEntry.js',
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
                    }
                }
            })
        );
        return config;
    }
);