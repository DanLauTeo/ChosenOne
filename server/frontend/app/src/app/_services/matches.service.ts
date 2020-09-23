import { HttpClient } from  '@angular/common/http';
import { map } from  'rxjs/operators';
import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class MatchesService {

  basePath: string = "/api/v1"

  constructor(private httpClient: HttpClient) { }

  public getMatches(): Observable<string[]> {
    return this.httpClient.get<MatchesResponse>(`${this.basePath}/matches/`)
      .pipe(map(resp => resp.user_ids));
  }
}

class MatchesResponse {
    user_ids: string[];
}
