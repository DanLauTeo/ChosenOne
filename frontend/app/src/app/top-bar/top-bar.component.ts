import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AccountService } from '../_services/account.service';

@Component({
  selector: 'app-top-bar',
  templateUrl: './top-bar.component.html',
  styleUrls: ['./top-bar.component.css']
})
export class TopBarComponent implements OnInit {
  constructor(    
    private router: Router,
    private accountService: AccountService) { }

  ngOnInit(): void {
  }

  feed() {
    this.router.navigate(['/feed']);
  }

  chats() {
    this.router.navigate(['/chats']);
  }

  getmatched() {
    this.router.navigate(['/get-matched']);
  }
  
  profile(){
    this.router.navigate(['/profile']);    
  }
}
