'use strict';

const { resolve } = require('path');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const { DefinePlugin, NormalModuleReplacementPlugin } = require('webpack');
const { BundleAnalyzerPlugin } = require('webpack-bundle-analyzer');
const NodemonPlugin = require('nodemon-webpack-plugin');
const packageJson = require('./package.json');

module.exports = (env = {}) => {
  const config = {
    entry: ['./src/main.ts'],
    mode: env.development ? 'development' : 'production',
    target: 'node',
    // devtool alternatives: cheap-module-eval-source-map (faster, less details) or cheap-eval-source-map (fastest, even less details)
    devtool: env.development ? 'inline-source-map' : false,
    node: {
      __dirname: false, // Fix for native node __dirname
      __filename: false, // Fix for native node __filename
    },
    output: {
      filename: `${packageJson.name}.js`,
      path: resolve(__dirname, 'dist'),
    },
    externals: { // https://github.com/matthew-andrews/isomorphic-fetch/issues/194
      'node-fetch': 'commonjs2 node-fetch',
    },
    resolve: {
      extensions: ['.ts', '.js'],
      modules: ['node_modules', 'src'],
    },
    stats: {
      modules: false, // We don't need to see this
    },
    module: {
      rules: [
        {
          test: /\.ts$/,
          use: 'ts-loader',
          exclude: /node_modules/,
        },
        {
          test: /\.node$/,
          use: 'raw-loader', //@azure/identity/dist-esm/src/credentials/visualStudioCodeCredential.js -> keytar/lib/keytar.js -> keytar/build/Release/keytar.node
        }
      ],
    },
    plugins: [
      new CleanWebpackPlugin(),
      new DefinePlugin({
        VERSION: JSON.stringify(packageJson.version),
        DEVELOP: !!env.development,
      }),
      // Use module replacement to use different configs for dev and prod
      new NormalModuleReplacementPlugin(
        /[\\/]src[\\/]config[\\/]config.ts$/, // [\\/] works on all operating systems.
        env.development ? 'config.dev.ts' : 'config.ts'
      ),
    ],
  };

  if (env.nodemon) {
    config.watch = true;
    config.plugins.push(new NodemonPlugin());
  }

  if (env.analyse) {
    config.plugins.push(
      new BundleAnalyzerPlugin({
        analyzerMode: 'static', // Generates file instead of starting a web server
      })
    );
  }

  return config;
};
