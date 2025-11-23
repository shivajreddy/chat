import { Component, signal } from '@angular/core';
import {Toggle} from '../toggle'

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [Toggle],
    template: `
    <main class="main">
        <div class="content">
        <p>Chat</p>
        <p>parent-val: {{ is_checked() }}</p>
        <toggle
            [checked] = "is_checked()"
            (changed) = "toggle_check()">
        </toggle>
        </div>
    </main>
`,
    styles: `
    main {
        height: 100vh;
        background: black;
        color: white;
    }
`,
})

export class App {
    is_checked = signal(false);
    toggle_check(){
        this.is_checked.update(v => !v);
    }
}
