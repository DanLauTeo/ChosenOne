import { HttpClient, HttpEvent, HttpErrorResponse, HttpEventType, HttpParams } from  '@angular/common/http';  
import { map } from  'rxjs/operators';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ProfileService {
  routeURL: string ;

  constructor(private httpClient: HttpClient) { }

  public upload(id, formData) {
    this.routeURL = "http://localhost:8000/user/"+id+"/profile-image";
    return this.httpClient.put<any>(this.routeURL, formData, {
      reportProgress: true,  
      observe: 'events'  
    });
  }
}
