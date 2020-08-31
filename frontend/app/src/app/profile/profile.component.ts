import { Component, OnInit } from '@angular/core';
import { User } from '../_models/user';
import { AccountService } from '../_services/account.service'
import { Observable } from 'rxjs';
import { USER } from '../_mock_models/mock_user'
@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  user : User;

  constructor( private accountService : AccountService) {
    //this.accountService.user.subscribe(x => this.user = x);
    this.user = accountService.getUser();
  }

  ngOnInit(): void {
  }

  edit() {
    if(document.getElementById("username").contentEditable ==  "true"){
      document.getElementById("username").contentEditable = "false";
      document.getElementById("bio").contentEditable = "false";
    } else {
    document.getElementById("username").contentEditable = "true";
    document.getElementById("bio").contentEditable = "true";
    }
  }

  logout() {
    this.accountService.logout();
  }
}
