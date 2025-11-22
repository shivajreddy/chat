import { ApplicationConfig, provideBrowserGlobalErrorListeners } from "@angular/core";

export const cardConfig: ApplicationConfig = {
    providers: [
        provideBrowserGlobalErrorListeners(),
    ]
}
