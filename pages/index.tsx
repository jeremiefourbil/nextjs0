import { NextPage } from 'next';

const Home: NextPage<{ userAgent: string }> = ({ userAgent }) => (
  <div>
    <h1>Hello world!</h1>
    <h2>user agent: {userAgent}</h2>
  </div>
);

Home.getInitialProps = async ({ req }) => {
  const userAgent = req ? req.headers['user-agent'] || '' : navigator.userAgent;
  return { userAgent };
};

export default Home;