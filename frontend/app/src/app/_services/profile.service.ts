import { HttpClient } from  '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { User } from '../_models/user'

@Injectable({
  providedIn: 'root'
})
export class ProfileService {
  constructor(private httpClient: HttpClient) { }

  public getUser(id: string): Observable<User> {
    return this.httpClient.get<User>(`/user/${id}/`);
  }

  public patchUser(id, patch) {
    return this.httpClient.patch<any>(`/user/${id}/`, patch);
  }

  public uploadProfilePic(id, formData) {
    console.log(formData)
    return this.httpClient.put("user/"+id+"/profile-image/", formData);
  }
}
