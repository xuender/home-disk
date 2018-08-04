import { NgModule } from '@angular/core';
import { IonicModule } from 'ionic-angular'
import { FileUploadModule } from 'ng2-file-upload'
//import { NgxUploaderModule } from 'ngx-uploader'

import { HdUploadComponent } from './hd-upload/hd-upload';
import { HdFilesComponent } from './hd-files/hd-files';
@NgModule({
  declarations: [
    HdUploadComponent,
    HdFilesComponent,
  ],
  imports: [
    IonicModule,
    FileUploadModule,
  ],
  exports: [
    HdUploadComponent,
    FileUploadModule,
    HdFilesComponent,
  ]
})
export class ComponentsModule { }
