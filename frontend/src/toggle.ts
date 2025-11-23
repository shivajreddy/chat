import {Component} from '@angular/core'

@Component({
    selector: 'toggle',
    standalone: true,
    template: `
    <input id="theme_toggle_input" type="checkbox" checked/>
    <p>toggle?</p>
`
})

export class Toggle {}
