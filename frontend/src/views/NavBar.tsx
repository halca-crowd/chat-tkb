import logo from '@/assets/chatTKB.svg';
import icon from '@/assets/sohosai.svg';
import { css } from '@emotion/react';
import styled from '@emotion/styled';

const NavbarContainer = styled.nav`
  position: sticky; /* スクロール時に固定 */
  top: 0; /* 上端に固定 */
  width: 100%; /* 画面幅いっぱいに広がる */
  height: 100px; /* 高さを指定 */
  display: flex;
  background-color: #ffffff;
  padding: 10px;
  border-bottom: 4px solid #26B4C5; 
  align-items: center;

  @media (min-width: 768px) {
    flex-direction: row;
    height: 100px;
  }
`;

const iconStyle = css`
  width: 85px;
  height: 85px;
  padding: 5px;

  @media (min-width: 768px) {
    width: 90px; /* スマートフォン用の幅 */
    height: 90px; /* スマートフォン用の高さ */
  }
`;
const logoStyle = css`
  width: 180px;
  height: 90px;  
  justify-content: center;

  @media (min-width: 768px) {
    width: 200px; /* スマートフォン用の幅 */
    height: 95px; /* スマートフォン用の高さ */
  }
`;

const NavbarLeft = styled.div`
  display: flex;
  margin-right: 10px;
  height: 100%;
  align-items: center;
`;

const NavbarCenter = styled.div`
  flex: 1;
  height: 100%;
  align-items: center;

`;

export function NavBar() {
    return (
        <NavbarContainer>
            <NavbarLeft>
                <img src={icon} css={iconStyle} alt="icon" />
            </NavbarLeft>
            <NavbarCenter>
                <img src={logo} css={logoStyle} alt="logo" />
            </NavbarCenter>
        </NavbarContainer>
    )
}

export default NavBar