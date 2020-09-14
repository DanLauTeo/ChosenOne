import { Component, OnInit, Input } from '@angular/core';
import * as Rx from 'rxjs';
import { MatchesService } from '../_services/matches.service';
import {USERS} from '../_mock_models/mock_users'
import {User} from '../_models/user'

@Component({
  selector: 'app-get-matched',
  templateUrl: './get-matched.component.html',
  styleUrls: ['./get-matched.component.css']
})
export class GetMatchedComponent implements OnInit {

  private matchesService: MatchesService;

  //matches: string[];
  matches: User[];

  constructor(
    //matchesService: MatchesService
  ) {
    //this.matchesService = matchesService;
  }

  ngOnInit(): void {
    //this.matchesService.getMatches()
    //  .subscribe(next => this.matches = next)
    this.matches =  USERS
  }

}
