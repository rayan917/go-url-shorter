import { TestBed } from '@angular/core/testing';

import { PastebinService } from './services/pastebin.service';

describe('PastebinService', () => {
  let service: PastebinService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PastebinService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
