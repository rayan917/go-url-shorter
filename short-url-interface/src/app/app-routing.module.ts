import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import PastebinComponent from './pastebin/pastebin.component';
import { HttpClientModule } from '@angular/common/http';

const routes: Routes = [
  { path: '', redirectTo: 'pastebin', pathMatch: 'full' },
  { path: 'pastebin', component: PastebinComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes),HttpClientModule],
  exports: [RouterModule],
  providers: [ HttpClientModule]
})
export class AppRoutingModule { }
