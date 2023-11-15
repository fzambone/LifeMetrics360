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
                shared: require('./package.json').dependencies,
            })
        );
        return config;
    }
);