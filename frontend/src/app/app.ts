import { Component, signal, computed } from '@angular/core';
import {Toggle} from '../toggle'

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [Toggle],
    template: `
    <main class="main"
[class.dark]="is_checked()"
[class.light]="!is_checked()"
    >
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
    }
    main.dark {
        background: black;
        color: white;
    }
    main.light {
        background: white;
        color: black;
    }
`,

})

export class App {
    is_checked = signal(false);
    toggle_check(){
        this.is_checked.update(v => !v);
    }
}
