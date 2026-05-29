import React, {type ReactNode} from 'react';
import TOC from '@theme-original/TOC';
import type TOCType from '@theme/TOC';
import type {WrapperProps} from '@docusaurus/types';

type Props = WrapperProps<typeof TOCType>;

export default function TOCWrapper(props: Props): ReactNode {
  return (
    <>
      <TOC {...props} />
      <div style={{ marginTop: '1.5rem', display: 'flex', flexDirection: 'column', alignItems: 'center', gap: '0.75rem' }}>
        <iframe src="https://github.com/sponsors/balajz/button" title="Sponsor balajz" height="32" width="114" style={{ border: 0, borderRadius: '6px' }}></iframe>
        <a href="https://www.paypal.com/paypalme/BalajiJothi01" target="_blank" rel="noopener noreferrer">
          <img src="https://img.shields.io/badge/PayPal-003087?style=for-the-badge&logo=paypal&logoColor=white" alt="Donate with PayPal" style={{ borderRadius: '6px' }} />
        </a>
      </div>
    </>
  );
}
