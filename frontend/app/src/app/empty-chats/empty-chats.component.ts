import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-empty-chats',
  templateUrl: './empty-chats.component.html',
  styleUrls: ['./empty-chats.component.css']
})
export class EmptyChatsComponent implements OnInit {

  constructor(private router: Router,) { }

  ngOnInit(): void {
  }

  getmatched() {
    this.router.navigate(['/get-matched']);
  }

}
