import {Component, Input, Output, EventEmitter } from '@angular/core'

@Component({
    selector: 'toggle',
    standalone: true,
    template: `
    <input
    id="theme_toggle_input"
    type="checkbox"
    [checked] = "checked"
    (change) = "changed.emit()"
/>
`
})

export class Toggle {
    @Input() checked: boolean = false;
    @Output() changed = new EventEmitter<void>();
}
