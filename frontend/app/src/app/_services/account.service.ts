import { Injectable } from '@angular/core';
//import { HttpClient } from '@angular/common/http';
import { BehaviorSubject, Observable, of  } from 'rxjs';
import { User } from '../_models/user';
import { USER } from '../_mock_models/mock_user';

@Injectable({
  providedIn: 'root'
})
export class AccountService {

  constructor() { }

  getUser(): Observable<User> {
    return of(USER);
  }
}
