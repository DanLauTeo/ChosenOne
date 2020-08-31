import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { AccountService } from '../_services/account.service';

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
      if (this.accountService.getUser())
        this.router.navigate(['/profile']);
   }

  ngOnInit(): void {
  }
  
  join()  {
    this.accountService.setUser();
    this.router.navigate(['/profile']);
  }

}
