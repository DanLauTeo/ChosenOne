import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { AccountService } from '../_services/account.service';
import { filter, take } from 'rxjs/operators';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  submitted = false;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private accountService: AccountService,
  ) {
      // redirect to home if already logged in
      this.accountService.user
        .subscribe(_ => this.router.navigate(["/profile"]));
   }

  ngOnInit(): void {
  }

  join()  {
    this.accountService.didLogin();
  }

}
