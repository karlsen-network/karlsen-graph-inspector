import { packageVersion } from "../version";

export type AppConfig = {
    karlsendVersion: string,
    processingVersion: string,
    network: string,
    apiVersion: string,
    webVersion: string,
};

export function getDefaultAppConfig(): AppConfig {
    return {
        karlsendVersion: "n/a",
        processingVersion: "n/a",
        network: "n/a",
        apiVersion: "n/a",
        webVersion: packageVersion,
    };
}

export function areAppConfigsEqual(left: AppConfig, right: AppConfig): boolean {
    return left.karlsendVersion === right.karlsendVersion
        && left.processingVersion === right.processingVersion
        && left.network === right.network
        && left.apiVersion === right.apiVersion
        && left.webVersion === right.webVersion;
}

export function isTestnet(appConfig: AppConfig): boolean {
    return appConfig.network.startsWith("karlsen-testnet");
}

export function isMainnet(appConfig: AppConfig): boolean {
    return appConfig.network === "karlsen-mainnet";
}
