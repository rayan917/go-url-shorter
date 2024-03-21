export interface Link {
    key: string;
    value: string;
    click?:number;
  }
  
  export interface Result {
    total_links: number;
    links: Link[];
  }
  
  export interface ShortURL {
    url: string;
  }
  
  export interface Stats {
    count: number;
  }
  
  export interface URLData {
    url: string;
  }