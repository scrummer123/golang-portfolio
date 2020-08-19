const path = require("path");

const hwp = require("html-webpack-plugin");

const isDev = process.env.NODE_ENV === "development";

module.exports = {
    devtool: "hidden-source-map",
    mode: process.env.NODE_ENV,
    entry: "./src/app/index.tsx",
    devServer: {
        compress: true,
        stats: {
            colors: true,
            hash: false,
            version: false,
            timings: false,
            assets: true,
            chunks: false,
            modules: false,
            reasons: false,
            children: false,
            source: false,
            errors: true,
            errorDetails: false,
            warnings: true,
            publicPath: false,
            chunkModules: false,
            entrypoints: false,
        },
    },
    optimization: {
        splitChunks: {
            chunks: "all",
        },
    },
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                use: {
                    loader: "ts-loader",
                    options: {
                        silent: true,
                    },
                },
                exclude: /node_modules/,
            },
        ],
    },
    resolve: {
        extensions: [".ts", ".tsx", ".js"],
    },
    output: {
        filename: isDev ? "[name].bundle.js" : "[name].[hash].bundle.js",
        path: path.resolve(__dirname, "src", "build"),
    },
    plugins: [new hwp({ template: "./src/app/static/index.html" })],
};
