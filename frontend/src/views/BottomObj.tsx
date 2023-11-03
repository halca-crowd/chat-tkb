import { css } from '@emotion/react';

const bottomObjectStyle = css`
position: fixed;
bottom: 0;
left: 0;
width: 100%;
height: 8%;
background-color: #D9D9D9;
color: #434343;
padding: 20px;
text-align: center;
`;

export function BottomObj (){
  return (
    <div css={bottomObjectStyle}>
      雙峰祭 創基151年 筑波大学 開学50周年企画：ChatGPT時代の就職、仕事、働き方
    </div>
  );
};

export default BottomObj;