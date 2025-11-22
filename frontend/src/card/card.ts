import { Component, signal } from "@angular/core";

@Component({
    selector: 'card',   /* selectors are always kebab-case */
    imports: [],
    templateUrl: './card.html',
})

export class Card {
    protected readonly title = signal('card');
}
