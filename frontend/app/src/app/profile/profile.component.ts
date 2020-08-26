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
  user = USER;
  constructor( private accountService : AccountService) {
  }

  ngOnInit(): void {
  }
}
