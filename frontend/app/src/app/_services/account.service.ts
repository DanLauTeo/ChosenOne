import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject, Observable, of  } from 'rxjs';
import { Router } from '@angular/router';
import { environment } from '../../environments/environment';
import { map } from 'rxjs/operators';

import { User } from '../_models/user';
import { USER } from '../_mock_models/mock_user';

@Injectable({
  providedIn: 'root'
})
export class AccountService {
  user : User;
  constructor( private router: Router) {
   }

  getUser(): User {
    return this.user;
  }

  getUserID(): string {
    return this.user.id;
  }

  setUser(){
    this.user = USER;
  }

  logout(){
    this.user = null;
    this.router.navigate(['/login']);
  }

  /*
  private userSubject: BehaviorSubject<User>;
  public user: Observable<User>;

  constructor(
    private router: Router,
    private http: HttpClient
  ) {
      this.userSubject = new BehaviorSubject<User>(JSON.parse(localStorage.getItem('user')));
      this.user = this.userSubject.asObservable();
  }

  public get userValue(): User {
    return this.userSubject.value;
  }

  login(){
    return this.http.post<User>(`${environment.apiUrl}/users/authenticate`, { username })
    .pipe(map(user => {
        // store user details and jwt token in local storage to keep user logged in between page refreshes
        localStorage.setItem('user', JSON.stringify(user));
        this.userSubject.next(user);
        return user;
    }));
  }

  logout() {
    // remove user from local storage and set current user to null
    localStorage.removeItem('user');
    this.userSubject.next(null);
    this.router.navigate(['/account/login']);
  } 
  */
}
