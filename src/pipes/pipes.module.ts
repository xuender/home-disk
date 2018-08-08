import { NgModule } from '@angular/core';
import { DayPipe } from './day/day';
import { FileTypePipe } from './file-type/file-type';
@NgModule({
  declarations: [
    DayPipe,
    FileTypePipe,
  ],
  imports: [],
  exports: [
    DayPipe,
    FileTypePipe,
  ]
})
export class PipesModule { }
