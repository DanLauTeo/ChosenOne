import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject, Observable, of  } from 'rxjs';
import { Router } from '@angular/router';
import { map, delay, filter } from 'rxjs/operators';

import { User } from '../_models/user';

@Injectable({
  providedIn: 'root'
})
export class AccountService {

  basePath: string = "/api/v1"

  loggedIn: boolean = false;

  getUser(): User {
    return this.userSubject.value;
  }

  getUserID(): string {
    return this.getUser().id;
  }

  didLogin() {
    this.httpClient.get<User>(`${this.basePath}/user/me/`)
      .subscribe(user => this.userSubject.next(user));
  }

  didLogout() {
    this.userSubject.next(null);
  }

  private userSubject: BehaviorSubject<User>;
  public user: Observable<User>;

  constructor(
    private router: Router,
    private httpClient: HttpClient
  ) {
      this.userSubject = new BehaviorSubject<User>(null);
      this.user = this.userSubject.asObservable()
        .pipe(filter((user, _) => user != null));
      this.didLogin();
  }

}
