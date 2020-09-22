import { HttpClient } from  '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { User } from '../_models/user'

@Injectable({
  providedIn: 'root'
})
export class ProfileService {

  basePath: string = "/api/v1"

  constructor(private httpClient: HttpClient) { }

  public getUser(id: string): Observable<User> {
    return this.httpClient.get<User>(`${this.basePath}/user/${id}/`);
  }

  public patchUser(id, patch) {
    return this.httpClient.patch<any>(`${this.basePath}/user/${id}/`, patch);
  }

  public uploadProfilePic(id, formData) {
    return this.httpClient.put(`${this.basePath}/user/${id}/profile-image/`, formData);
  }
}
