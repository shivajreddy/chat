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
        <toggle></toggle>
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
  protected readonly title = signal('frontend');
}
